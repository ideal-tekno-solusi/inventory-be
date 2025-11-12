package id.my.idtecsi.inventory.api.inventory.handler;

import java.util.List;
import java.util.UUID;

import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import id.my.idtecsi.inventory.api.inventory.model.ErrorResponseDefault;
import id.my.idtecsi.inventory.api.inventory.model.ErrorValidation;
import id.my.idtecsi.inventory.api.inventory.model.GlobalIdentifiedException;
import jakarta.servlet.http.HttpServletRequest;

@RestControllerAdvice
public class ErrorController {
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ErrorResponseDefault<List<ErrorValidation>>> handleValidationExceptions(
            MethodArgumentNotValidException ex, HttpServletRequest request) {
        HttpHeaders headers = new HttpHeaders();
        headers.add("Content-Type", "application/problem+json");

        List<ErrorValidation> errors = ex.getBindingResult()
                .getFieldErrors()
                .stream()
                .map(err -> {
                    ErrorValidation ev = new ErrorValidation();
                    ev.setField(err.getField());
                    ev.setMessage(err.getDefaultMessage());
                    return ev;
                })
                .toList();

        ErrorResponseDefault<List<ErrorValidation>> errorMessage = new ErrorResponseDefault<List<ErrorValidation>>();
        errorMessage.setType(HttpStatus.BAD_REQUEST.name());
        errorMessage.setTitle("Bad Request");
        errorMessage.setStatus(HttpStatus.BAD_REQUEST.value());
        errorMessage.setMessage("One or more validation errors occurred.");
        errorMessage.setErrors(errors);
        errorMessage.setInstance(request.getRequestURI());
        errorMessage.setGuid(UUID.randomUUID().toString());

        return new ResponseEntity<ErrorResponseDefault<List<ErrorValidation>>>(errorMessage, headers,
                HttpStatus.BAD_REQUEST);
    }

    @ExceptionHandler(GlobalIdentifiedException.class)
    public ResponseEntity<ErrorResponseDefault<String>> handleGlobalIdentifiedExceptions(
            GlobalIdentifiedException ex, HttpServletRequest request) {
        HttpHeaders headers = new HttpHeaders();
        headers.add("Content-Type", "application/problem+json");

        ErrorResponseDefault<String> errorMessage = new ErrorResponseDefault<String>();
        errorMessage.setType(HttpStatus.valueOf(ex.getStatus()).name());
        errorMessage.setTitle(HttpStatus.valueOf(ex.getStatus()).getReasonPhrase());
        errorMessage.setStatus(ex.getStatus());
        errorMessage.setMessage(ex.getMessage());
        errorMessage.setInstance(request.getRequestURI());
        errorMessage.setGuid(UUID.randomUUID().toString());

        return new ResponseEntity<ErrorResponseDefault<String>>(errorMessage, headers,
                HttpStatus.valueOf(ex.getStatus()));
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponseDefault<String>> handleGenericExceptions(
            Exception ex, HttpServletRequest request) {
        HttpHeaders headers = new HttpHeaders();
        headers.add("Content-Type", "application/problem+json");

        ErrorResponseDefault<String> errorMessage = new ErrorResponseDefault<String>();
        errorMessage.setType(HttpStatus.INTERNAL_SERVER_ERROR.name());
        errorMessage.setTitle("Internal Server Error");
        errorMessage.setStatus(HttpStatus.INTERNAL_SERVER_ERROR.value());
        errorMessage.setMessage(ex.getMessage());
        errorMessage.setInstance(request.getRequestURI());
        errorMessage.setGuid(UUID.randomUUID().toString());

        return new ResponseEntity<ErrorResponseDefault<String>>(errorMessage, headers,
                HttpStatus.INTERNAL_SERVER_ERROR);
    }
}
