package id.my.idtecsi.inventory.api.inventory.repository;

import java.sql.SQLException;
import java.util.List;

import id.my.idtecsi.inventory.api.inventory.entity.Inventory;
import id.my.idtecsi.inventory.api.inventory.entity.PageInformation;

public interface InventoryService {
    PageInformation getInventoryPageInformation(String categoryId, String BranchId, int limit) throws SQLException;

    List<Inventory> getInventoryList(String categoryId, String BranchId, int page, int limit) throws SQLException;
}
