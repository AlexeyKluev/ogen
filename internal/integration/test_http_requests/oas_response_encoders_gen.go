// Code generated by ogen, DO NOT EDIT.

package api

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeAllRequestBodiesResponse(response AllRequestBodiesOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeAllRequestBodiesOptionalResponse(response AllRequestBodiesOptionalOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := w
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeBase64RequestResponse(response Base64RequestOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	writer := base64.NewEncoder(base64.StdEncoding, w)
	defer writer.Close()
	if _, err := io.Copy(writer, response); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeMaskContentTypeResponse(response *MaskResponse, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeMaskContentTypeOptionalResponse(response *MaskResponse, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeStreamJSONResponse(response float64, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	e.Float64(response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}
