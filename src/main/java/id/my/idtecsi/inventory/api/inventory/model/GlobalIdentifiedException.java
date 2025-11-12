package id.my.idtecsi.inventory.api.inventory.model;

public class GlobalIdentifiedException extends Exception {
    private int status;
    private String message;

    public GlobalIdentifiedException(int status, String message) {
        super(message);
        this.status = status;
        this.message = message;
    }

    public int getStatus() {
        return status;
    }

    public String getMessage() {
        return message;
    }
}
