package apiresponse

import "net/http"

// Accepted writes a response to the Responder with the status of http.StatusAccepted.
func (r Responder) Accepted() { r.write(http.StatusAccepted) }

// AlreadyReported writes a response to the Responder with the status of http.StatusAlreadyReported.
func (r Responder) AlreadyReported() { r.write(http.StatusAlreadyReported) }

// BadGateway writes a response to the Responder with the status of http.StatusBadGateway.
func (r Responder) BadGateway() { r.write(http.StatusBadGateway) }

// BadRequest writes a response to the Responder with the status of http.StatusBadRequest.
func (r Responder) BadRequest() { r.write(http.StatusBadRequest) }

// Conflict writes a response to the Responder with the status of http.StatusConflict.
func (r Responder) Conflict() { r.write(http.StatusConflict) }

// Continue writes a response to the Responder with the status of http.StatusContinue.
func (r Responder) Continue() { r.write(http.StatusContinue) }

// Created writes a response to the Responder with the status of http.StatusCreated.
func (r Responder) Created() { r.write(http.StatusCreated) }

// ExpectationFailed writes a response to the Responder with the status of http.StatusExpectationFailed.
func (r Responder) ExpectationFailed() { r.write(http.StatusExpectationFailed) }

// FailedDependency writes a response to the Responder with the status of http.StatusFailedDependency.
func (r Responder) FailedDependency() { r.write(http.StatusFailedDependency) }

// Forbidden writes a response to the Responder with the status of http.StatusForbidden.
func (r Responder) Forbidden() { r.write(http.StatusForbidden) }

// Found writes a response to the Responder with the status of http.StatusFound.
func (r Responder) Found() { r.write(http.StatusFound) }

// GatewayTimeout writes a response to the Responder with the status of http.StatusGatewayTimeout.
func (r Responder) GatewayTimeout() { r.write(http.StatusGatewayTimeout) }

// Gone writes a response to the Responder with the status of http.StatusGone.
func (r Responder) Gone() { r.write(http.StatusGone) }

// HTTPVersionNotSupported writes a response to the Responder with the status of http.StatusHTTPVersionNotSupported.
func (r Responder) HTTPVersionNotSupported() { r.write(http.StatusHTTPVersionNotSupported) }

// IMUsed writes a response to the Responder with the status of http.StatusIMUsed.
func (r Responder) IMUsed() { r.write(http.StatusIMUsed) }

// InsufficientStorage writes a response to the Responder with the status of http.StatusInsufficientStorage.
func (r Responder) InsufficientStorage() { r.write(http.StatusInsufficientStorage) }

// InternalServerError writes a response to the Responder with the status of http.StatusInternalServerError.
func (r Responder) InternalServerError() { r.write(http.StatusInternalServerError) }

// LengthRequired writes a response to the Responder with the status of http.StatusLengthRequired.
func (r Responder) LengthRequired() { r.write(http.StatusLengthRequired) }

// Locked writes a response to the Responder with the status of http.StatusLocked.
func (r Responder) Locked() { r.write(http.StatusLocked) }

// LoopDetected writes a response to the Responder with the status of http.StatusLoopDetected.
func (r Responder) LoopDetected() { r.write(http.StatusLoopDetected) }

// MethodNotAllowed writes a response to the Responder with the status of http.StatusMethodNotAllowed.
func (r Responder) MethodNotAllowed() { r.write(http.StatusMethodNotAllowed) }

// MisdirectedRequest writes a response to the Responder with the status of http.StatusMisdirectedRequest.
func (r Responder) MisdirectedRequest() { r.write(http.StatusMisdirectedRequest) }

// MovedPermanently writes a response to the Responder with the status of http.StatusMovedPermanently.
func (r Responder) MovedPermanently() { r.write(http.StatusMovedPermanently) }

// MultiStatus writes a response to the Responder with the status of http.StatusMultiStatus.
func (r Responder) MultiStatus() { r.write(http.StatusMultiStatus) }

// MultipleChoices writes a response to the Responder with the status of http.StatusMultipleChoices.
func (r Responder) MultipleChoices() { r.write(http.StatusMultipleChoices) }

// NetworkAuthenticationRequired writes a response to the Responder with the status of http.StatusNetworkAuthenticationRequired.
func (r Responder) NetworkAuthenticationRequired() { r.write(http.StatusNetworkAuthenticationRequired) }

// NoContent writes a response to the Responder with the status of http.StatusNoContent.
func (r Responder) NoContent() { r.write(http.StatusNoContent) }

// NonAuthoritativeInfo writes a response to the Responder with the status of http.StatusNonAuthoritativeInfo.
func (r Responder) NonAuthoritativeInfo() { r.write(http.StatusNonAuthoritativeInfo) }

// NotAcceptable writes a response to the Responder with the status of http.StatusNotAcceptable.
func (r Responder) NotAcceptable() { r.write(http.StatusNotAcceptable) }

// NotExtended writes a response to the Responder with the status of http.StatusNotExtended.
func (r Responder) NotExtended() { r.write(http.StatusNotExtended) }

// NotFound writes a response to the Responder with the status of http.StatusNotFound.
func (r Responder) NotFound() { r.write(http.StatusNotFound) }

// NotImplemented writes a response to the Responder with the status of http.StatusNotImplemented.
func (r Responder) NotImplemented() { r.write(http.StatusNotImplemented) }

// NotModified writes a response to the Responder with the status of http.StatusNotModified.
func (r Responder) NotModified() { r.write(http.StatusNotModified) }

// OK writes a response to the Responder with the status of http.StatusOK.
func (r Responder) OK() { r.write(http.StatusOK) }

// PartialContent writes a response to the Responder with the status of http.StatusPartialContent.
func (r Responder) PartialContent() { r.write(http.StatusPartialContent) }

// PaymentRequired writes a response to the Responder with the status of http.StatusPaymentRequired.
func (r Responder) PaymentRequired() { r.write(http.StatusPaymentRequired) }

// PermanentRedirect writes a response to the Responder with the status of http.StatusPermanentRedirect.
func (r Responder) PermanentRedirect() { r.write(http.StatusPermanentRedirect) }

// PreconditionFailed writes a response to the Responder with the status of http.StatusPreconditionFailed.
func (r Responder) PreconditionFailed() { r.write(http.StatusPreconditionFailed) }

// PreconditionRequired writes a response to the Responder with the status of http.StatusPreconditionRequired.
func (r Responder) PreconditionRequired() { r.write(http.StatusPreconditionRequired) }

// Processing writes a response to the Responder with the status of http.StatusProcessing.
func (r Responder) Processing() { r.write(http.StatusProcessing) }

// ProxyAuthRequired writes a response to the Responder with the status of http.StatusProxyAuthRequired.
func (r Responder) ProxyAuthRequired() { r.write(http.StatusProxyAuthRequired) }

// RequestEntityTooLarge writes a response to the Responder with the status of http.StatusRequestEntityTooLarge.
func (r Responder) RequestEntityTooLarge() { r.write(http.StatusRequestEntityTooLarge) }

// RequestHeaderFieldsTooLarge writes a response to the Responder with the status of http.StatusRequestHeaderFieldsTooLarge.
func (r Responder) RequestHeaderFieldsTooLarge() { r.write(http.StatusRequestHeaderFieldsTooLarge) }

// RequestTimeout writes a response to the Responder with the status of http.StatusRequestTimeout.
func (r Responder) RequestTimeout() { r.write(http.StatusRequestTimeout) }

// RequestURITooLong writes a response to the Responder with the status of http.StatusRequestURITooLong.
func (r Responder) RequestURITooLong() { r.write(http.StatusRequestURITooLong) }

// RequestedRangeNotSatisfiable writes a response to the Responder with the status of http.StatusRequestedRangeNotSatisfiable.
func (r Responder) RequestedRangeNotSatisfiable() { r.write(http.StatusRequestedRangeNotSatisfiable) }

// ResetContent writes a response to the Responder with the status of http.StatusResetContent.
func (r Responder) ResetContent() { r.write(http.StatusResetContent) }

// SeeOther writes a response to the Responder with the status of http.StatusSeeOther.
func (r Responder) SeeOther() { r.write(http.StatusSeeOther) }

// ServiceUnavailable writes a response to the Responder with the status of http.StatusServiceUnavailable.
func (r Responder) ServiceUnavailable() { r.write(http.StatusServiceUnavailable) }

// SwitchingProtocols writes a response to the Responder with the status of http.StatusSwitchingProtocols.
func (r Responder) SwitchingProtocols() { r.write(http.StatusSwitchingProtocols) }

// Teapot writes a response to the Responder with the status of http.StatusTeapot.
func (r Responder) Teapot() { r.write(http.StatusTeapot) }

// TemporaryRedirect writes a response to the Responder with the status of http.StatusTemporaryRedirect.
func (r Responder) TemporaryRedirect() { r.write(http.StatusTemporaryRedirect) }

// TooManyRequests writes a response to the Responder with the status of http.StatusTooManyRequests.
func (r Responder) TooManyRequests() { r.write(http.StatusTooManyRequests) }

// Unauthorized writes a response to the Responder with the status of http.StatusUnauthorized.
func (r Responder) Unauthorized() { r.write(http.StatusUnauthorized) }

// UnavailableForLegalReasons writes a response to the Responder with the status of http.StatusUnavailableForLegalReasons.
func (r Responder) UnavailableForLegalReasons() { r.write(http.StatusUnavailableForLegalReasons) }

// UnprocessableEntity writes a response to the Responder with the status of http.StatusUnprocessableEntity.
func (r Responder) UnprocessableEntity() { r.write(http.StatusUnprocessableEntity) }

// UnsupportedMediaType writes a response to the Responder with the status of http.StatusUnsupportedMediaType.
func (r Responder) UnsupportedMediaType() { r.write(http.StatusUnsupportedMediaType) }

// UpgradeRequired writes a response to the Responder with the status of http.StatusUpgradeRequired.
func (r Responder) UpgradeRequired() { r.write(http.StatusUpgradeRequired) }

// UseProxy writes a response to the Responder with the status of http.StatusUseProxy.
func (r Responder) UseProxy() { r.write(http.StatusUseProxy) }

// VariantAlsoNegotiates writes a response to the Responder with the status of http.StatusVariantAlsoNegotiates.
func (r Responder) VariantAlsoNegotiates() { r.write(http.StatusVariantAlsoNegotiates) }
