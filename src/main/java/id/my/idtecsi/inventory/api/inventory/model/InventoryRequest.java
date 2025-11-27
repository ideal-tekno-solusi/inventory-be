package id.my.idtecsi.inventory.api.inventory.model;

import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;

public class InventoryRequest {
    private String categoryId;
    private String branchId;
    @Min(value = 1, message = "page minimum is 1")
    private int page;
    @Max(value = 50, message = "limit maximum is 50")
    private int limit;

    public String getCategoryId() {
        return categoryId;
    }

    public void setCategoryId(String categoryId) {
        this.categoryId = categoryId;
    }

    public String getBranchId() {
        return branchId;
    }

    public void setBranchId(String branchId) {
        this.branchId = branchId;
    }

    public int getPage() {
        return page;
    }

    public void setPage(int page) {
        this.page = page;
    }

    public int getLimit() {
        return limit;
    }

    public void setLimit(int limit) {
        this.limit = limit;
    }
}
