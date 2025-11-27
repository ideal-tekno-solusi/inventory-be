package id.my.idtecsi.inventory.api.inventory.repository;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;

import id.my.idtecsi.inventory.api.inventory.bootstrap.DatabaseService;
import id.my.idtecsi.inventory.api.inventory.entity.Inventory;
import id.my.idtecsi.inventory.api.inventory.entity.PageInformation;

public class InventoryRepository implements InventoryService {
    private final Connection dbr;
    private final Connection dbw;

    public InventoryRepository(DatabaseService dbr, DatabaseService dbw) {
        this.dbr = dbr.getDb();
        this.dbw = dbw.getDb();
    }

    @Override
    public PageInformation getInventoryPageInformation(String categoryId, String BranchId, int limit)
            throws SQLException {
        PageInformation pageInfo = new PageInformation();

        String sql = """
                        select
                            count(*)
                        from
                            branch_items
                        join
                            items
                        on
                            branch_items.item_id = items.id
                        join
                            categories
                        on
                            items.category_id = categories.id
                        join
                            branches
                        on
                            branch_items.branch_id = branches.id
                        join
                            positions
                        on
                            branch_items.position_id = positions.id
                        where
                            categories.id ilike ?
                        and
                            branches.id ilike ?
                        and
                            branch_items.delete_date is null
                """;

        PreparedStatement stmt = dbr.prepareStatement(sql);
        stmt.setString(1, String.format("%%%s%%", categoryId));
        stmt.setString(2, String.format("%%%s%%", BranchId));

        ResultSet rs = stmt.executeQuery();

        while (rs.next()) {
            pageInfo.setTotalItems(rs.getInt(1));
            pageInfo.setTotalPages((int) Math.ceil((double) pageInfo.getTotalItems() / limit));
        }

        return pageInfo;
    }

    @Override
    public List<Inventory> getInventoryList(String categoryId, String BranchId, int page, int limit)
            throws SQLException {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'getInventoryList'");
    }

}
