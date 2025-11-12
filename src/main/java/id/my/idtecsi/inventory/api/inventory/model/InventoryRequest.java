package id.my.idtecsi.inventory.api.inventory.model;

import jakarta.validation.constraints.Max;

public class InventoryRequest {
    private String category;
    private String branchId;
    private int page;
    @Max(value = 50, message = "limit maximum is 50")
    private int limit;

    public String getCategory() {
        return category;
    }

    public void setCategory(String category) {
        this.category = category;
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
