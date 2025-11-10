package id.my.idtecsi.inventory.api.inventory.bootstrap;

import java.sql.Connection;
import java.sql.DriverManager;
import java.util.Properties;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.DisposableBean;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class PostgresDatabaseWriteService implements DatabaseService, DisposableBean {
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

    private static final Logger log = LoggerFactory.getLogger(PostgresDatabaseWriteService.class);

    @Override
    public Connection getDb() {
        if (dbw == null) {
            try {
                String dsn = "jdbc:postgresql://" + this.dbwHost + ":" + this.dbwPort + "/" + this.dbwDatabase;
                Properties props = new Properties();
                props.setProperty("user", this.dbwUsername);
                props.setProperty("password", this.dbwPassword);
                props.setProperty("currentSchema", this.dbwSchema);
                props.setProperty("ssl", "disable");

                Connection conn = DriverManager.getConnection(dsn, props);

                log.info("success connect to dbw");

                dbw = conn;
            } catch (Exception ex) {
                ex.printStackTrace();
                log.error("failed to connect to db with error: ", ex);
                System.exit(0);
            }
        }

        return dbw;
    }

    @Override
    public void close() {
        try {
            if (dbw != null) {
                dbw.close();
                log.info("success close dbw");
            }
        } catch (Exception ex) {
            log.error("failed to close connection with error: ", ex);
        }
    }

    @Override
    public void destroy() throws Exception {
        close();
    }
}
