package id.my.idtecsi.inventory.api.inventory.handler;

import java.util.List;
import java.util.UUID;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;

import id.my.idtecsi.inventory.api.inventory.entity.DomainErrorResponseDefault;
import id.my.idtecsi.inventory.api.inventory.model.ErrorValidation;
import jakarta.servlet.http.HttpServletRequest;

@ControllerAdvice
public class ErrorController {
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<DomainErrorResponseDefault<List<ErrorValidation>>> handleValidationExceptions(
            MethodArgumentNotValidException ex, HttpServletRequest request) {
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

        DomainErrorResponseDefault<List<ErrorValidation>> errorMessage = new DomainErrorResponseDefault<List<ErrorValidation>>();
        errorMessage.setErrors(errors);
        errorMessage.setStatus(HttpStatus.BAD_REQUEST.value());
        errorMessage.setType(HttpStatus.BAD_REQUEST.name());
        errorMessage.setTitle("Bad Request");
        errorMessage.setInstance(request.getRequestURI());
        errorMessage.setMessage("One or more validation errors occurred.");
        errorMessage.setGuid(UUID.randomUUID().toString());

        return new ResponseEntity<DomainErrorResponseDefault<List<ErrorValidation>>>(errorMessage,
                HttpStatus.BAD_REQUEST);
    }
}
