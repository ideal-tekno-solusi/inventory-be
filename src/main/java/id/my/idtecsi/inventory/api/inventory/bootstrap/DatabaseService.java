package id.my.idtecsi.inventory.api.inventory.bootstrap;

import java.sql.Connection;

public interface DatabaseService {
    Connection getDb();

    void close();
}
