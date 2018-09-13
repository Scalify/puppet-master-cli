package format

import (
	"fmt"
	"text/tabwriter"

	"github.com/Scalify/puppet-master-client-go"
	"github.com/sirupsen/logrus"
)

// Job formats a puppet master job to the target of given logger
// nolint: errcheck, gosec
func Job(logger *logrus.Entry, job *puppetmaster.Job) {
	fmt.Fprint(logger.Logger.Out, "\n\nJob:\n")

	w := tabwriter.NewWriter(logger.Logger.Out, 20, 10, 1, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "UUID\t%s\t\n", job.UUID)
	fmt.Fprintf(w, "Status\t%s\t\n", job.Status)
	fmt.Fprintf(w, "Duration\t%d\t\n", job.Duration)
	fmt.Fprintf(w, "Error\t%s\t\n", job.Error)
	fmt.Fprintf(w, "Created at\t%s\t\n", job.CreatedAt)
	fmt.Fprintf(w, "Started at\t%s\t\n", job.StartedAt)
	fmt.Fprintf(w, "Finished at\t%s\t\n", job.FinishedAt)
	w.Flush()

	fmt.Fprint(logger.Logger.Out, "\n\nLogs:\n")
	logs(logger, job.Logs)

	fmt.Fprint(logger.Logger.Out, "\n\nResults:\n")
	results(logger, job.Results)

	fmt.Fprint(logger.Logger.Out, "\n\n")
}

// nolint: errcheck, gosec
func logs(logger *logrus.Entry, logs []puppetmaster.Log) {
	w := tabwriter.NewWriter(logger.Logger.Out, 20, 10, 1, ' ', tabwriter.Debug)
	for _, l := range logs {
		fmt.Fprintf(w, "%s\t%s\t%s\t\n", l.Time, l.Level, l.Message)
	}
	w.Flush()
}

// nolint: errcheck, gosec
func results(logger *logrus.Entry, res map[string]interface{}) {
	w := tabwriter.NewWriter(logger.Logger.Out, 20, 10, 1, ' ', tabwriter.Debug)
	for k, v := range res {
		fmt.Fprintf(w, "%s\t%v\t\n", k, v)
	}
	w.Flush()
}
