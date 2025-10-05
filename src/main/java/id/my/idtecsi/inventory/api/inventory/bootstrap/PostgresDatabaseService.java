package id.my.idtecsi.inventory.api.inventory.bootstrap;

import java.sql.Connection;
import java.sql.DriverManager;
import java.util.Properties;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class PostgresDatabaseService implements DatabaseService {
    @Value("${database.read.host}")
    private String dbrHost;
    @Value("${database.read.port}")
    private String dbrPort;
    @Value("${database.read.database}")
    private String dbrDatabase;
    @Value("${database.read.schema}")
    private String dbrSchema;
    @Value("${database.read.username}")
    private String dbrUsername;
    @Value("${database.read.password}")
    private String dbrPassword;

    @Value("${database.write.host}")
    private String dbwHost;
    @Value("${database.write.port}")
    private String dbwPort;
    @Value("${database.write.database}")
    private String dbwDatabase;
    @Value("${database.write.schema}")
    private String dbwSchema;
    @Value("${database.write.username}")
    private String dbwUsername;
    @Value("${database.write.password}")
    private String dbwPassword;

    private static Connection dbw;
    private static Connection dbr;

    @Override
    public Connection getDbw() {
        if (dbw == null) {
            try {
                String dsn = "jdbc:postgresql://" + this.dbwHost + ":" + this.dbwPort + "/" + this.dbwDatabase;
                Properties props = new Properties();
                props.setProperty("user", this.dbwUsername);
                props.setProperty("password", this.dbwPassword);
                props.setProperty("currentSchema", this.dbwSchema);
                props.setProperty("ssl", "disable");

                Connection conn = DriverManager.getConnection(dsn, props);

                dbw = conn;
            } catch (Exception ex) {
                ex.printStackTrace();
                System.out.println(String.format("failed to connect to db with error: %s", ex.getMessage()));
                System.exit(0);
            }
        }

        return dbw;
    }

    @Override
    public Connection getDbr() {
        if (dbr == null) {
            try {
                String dsn = "jdbc:postgresql://" + this.dbrHost + ":" + this.dbrPort + "/" + this.dbrDatabase;
                Properties props = new Properties();
                props.setProperty("user", this.dbrUsername);
                props.setProperty("password", this.dbrPassword);
                props.setProperty("currentSchema", this.dbrSchema);
                props.setProperty("ssl", "disable");

                Connection conn = DriverManager.getConnection(dsn, props);

                dbr = conn;
            } catch (Exception ex) {
                ex.printStackTrace();
                System.out.println(String.format("failed to connect to db with error: %s", ex.getMessage()));
                System.exit(0);
            }
        }

        return dbr;
    }

    @Override
    public void Close() {
        try {
            if (dbr != null) {
                dbr.close();
                System.out.println("success close dbr");
            }

            if (dbw != null) {
                dbw.close();
                System.out.println("success close dbw");
            }
        } catch (Exception ex) {
            System.out.println("failed to close connection");
        }
    }
}
