package id.my.idtecsi.inventory.api.inventory.model;

public class InventoryResponse {
    private Inventory[] inventories;
    private Page page;

    public Inventory[] getInventories() {
        return inventories;
    }

    public void setInventories(Inventory[] inventories) {
        this.inventories = inventories;
    }

    public Page getPage() {
        return page;
    }

    public void setPage(Page page) {
        this.page = page;
    }
}
