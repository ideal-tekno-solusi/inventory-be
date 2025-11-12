package id.my.idtecsi.inventory.api.inventory.entity;

import java.time.LocalDateTime;

public class Response<T> {
    private String guid;
    private LocalDateTime timestamp;
    private boolean isSuccess;
    private T data;
}
