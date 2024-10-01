package plotter

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"

	"eth_usd/csv_process"
)

func GeneratePlotFor(pricePairs []csv_process.EthereumPrice) (*plot.Plot, error) {
	points := make(plotter.XYs, len(pricePairs))

	for i := range pricePairs {
		points[i].X = float64(pricePairs[i].Date.Unix()) // conver time to float64
		points[i].Y = pricePairs[i].Price
	}

	pricePlot := plot.New()

	pricePlot.Title.Text = "Etherium Price Chart"
	pricePlot.X.Label.Text = "Date"
	pricePlot.Y.Label.Text = "Price (USD)"

	if err := plotutil.AddLinePoints(pricePlot, "ETH", points); err != nil {
		return nil, fmt.Errorf("error adding line to plot: %w", err)
	}

	pricePlot.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}

	return pricePlot, nil
}
