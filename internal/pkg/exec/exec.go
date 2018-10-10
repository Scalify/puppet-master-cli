package exec

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Scalify/puppet-master-cli/internal/pkg/file"
	"github.com/Scalify/puppet-master-cli/internal/pkg/format"
	"github.com/Scalify/puppet-master-client-go"
	"github.com/sirupsen/logrus"
)

// Execute a job and print the finished one to loggers target
func Execute(logger *logrus.Entry, client *puppetmaster.Client, baseDir, codeFile, varsFile string, moduleFiles []string, delete, execLogsVerbose bool) error {
	code, err := file.Load(filepath.Join(baseDir, codeFile))
	if err != nil {
		return err
	}

	vars := make(map[string]string)
	if err = file.LoadJSON(filepath.Join(baseDir, varsFile), &vars); err != nil {
		return err
	}

	modules, err := loadModules(logger, baseDir, &moduleFiles)
	if err != nil {
		return err
	}

	logger.Infof("Loaded code file %s, vars file %s, module files %s", codeFile, varsFile, moduleFiles)

	job, err := client.CreateJob(&puppetmaster.JobRequest{
		Code:    code,
		Vars:    vars,
		Modules: modules,
	})
	if err != nil {
		return fmt.Errorf("failed to execute job: %v", err)
	}

	logger.Infof("Created job with id %s", job.UUID)
	for {
		job, err = client.GetJob(job.UUID)
		if err != nil {
			return fmt.Errorf("failed to get job: %v", err)
		}

		logger.Infof("Job has status %s", job.Status)
		if job.Status == puppetmaster.StatusDone {
			break
		}

		time.Sleep(time.Second)
	}

	format.Job(logger, job, execLogsVerbose)

	if delete {
		if err := client.DeleteJob(job.UUID); err != nil {
			return fmt.Errorf("failed to delete job: %v", err)
		}
	}

	return nil
}

func loadModules(logger *logrus.Entry, baseDir string, moduleFiles *[]string) (map[string]string, error) {
	modules := make(map[string]string)

	if len(*moduleFiles) == 0 {
		files, err := ioutil.ReadDir(filepath.Join(baseDir, "modules"))
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			return modules, err
		}

		for _, f := range files {
			if f.IsDir() || filepath.Ext(f.Name()) != ".mjs" {
				continue
			}

			*moduleFiles = append(*moduleFiles, filepath.Join("modules", f.Name()))
		}
	}

	var err error
	for _, m := range *moduleFiles {
		name := strings.Replace(filepath.Base(m), filepath.Ext(m), "", -1)
		modules[name], err = file.Load(filepath.Join(baseDir, m))
		if err != nil {
			return modules, err
		}
	}

	return modules, nil
}
