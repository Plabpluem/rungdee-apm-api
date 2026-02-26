package pdf

import (
	"bytes"
	"context"
	"html/template"
	"os"
	"path/filepath"
	usecases "rungdee-apm-api/internal/usecases/invoice"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type ChromePdfGenerate struct {
}

func NewChromePdfGenerate() usecases.PdfGenerate {
	return &ChromePdfGenerate{}
}

func (p *ChromePdfGenerate) Generate(data *usecases.PdfData) ([]byte, error) {
	htmlPath := filepath.Join("internal", "adapters", "invoice", "template", "invoice_template.html")
	templateContent, err := os.ReadFile(htmlPath)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("invoice").Parse(string(templateContent))
	if err != nil {
		return nil, err
	}

	var htmlBuf bytes.Buffer
	if err := tmpl.Execute(&htmlBuf, data); err != nil {
		return nil, err
	}
	// แปลงเป็นไฟล์ temp + navigate
	tmpFile, err := os.CreateTemp("", "invoice-*.html")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(htmlBuf.String()); err != nil {
		return nil, err
	}
	tmpFile.Close()

	fileUrl := "file://" + filepath.ToSlash(tmpFile.Name())

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var pdfBytes []byte
	err = chromedp.Run(ctx,
		chromedp.Navigate(fileUrl),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPaperWidth(8.27). // A4
				WithPaperHeight(11.69).
				WithMarginTop(0.4).
				WithMarginBottom(0.4).
				WithMarginLeft(0.4).
				WithMarginRight(0.4).
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				return err
			}
			pdfBytes = buf
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return pdfBytes, nil
}
