package id.my.idtecsi.inventory.api.inventory.model;

import java.time.LocalDateTime;

public class ResponseDefault<T> {
    private String guid;
    private LocalDateTime timestamp;
    private boolean isSuccess;
    private T data;

    public String getGuid() {
        return guid;
    }

    public void setGuid(String guid) {
        this.guid = guid;
    }

    public LocalDateTime getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(LocalDateTime timestamp) {
        this.timestamp = timestamp;
    }

    public boolean isSuccess() {
        return isSuccess;
    }

    public void setSuccess(boolean success) {
        isSuccess = success;
    }

    public T getData() {
        return data;
    }

    public void setData(T data) {
        this.data = data;
    }
}
