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
public class PostgresDatabaseReadService implements DatabaseService, DisposableBean {
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

    private static Connection dbr;

    private static final Logger log = LoggerFactory.getLogger(PostgresDatabaseReadService.class);

    @Override
    public Connection getDb() {
        if (dbr == null) {
            try {
                String dsn = "jdbc:postgresql://" + this.dbrHost + ":" + this.dbrPort + "/" + this.dbrDatabase;
                Properties props = new Properties();
                props.setProperty("user", this.dbrUsername);
                props.setProperty("password", this.dbrPassword);
                props.setProperty("currentSchema", this.dbrSchema);
                props.setProperty("ssl", "disable");

                Connection conn = DriverManager.getConnection(dsn, props);

                log.info("success connect to dbr");

                dbr = conn;
            } catch (Exception ex) {
                ex.printStackTrace();
                log.error("failed to connect to db with error: ", ex);
                System.exit(0);
            }
        }

        return dbr;
    }

    @Override
    public void close() {
        try {
            if (dbr != null) {
                dbr.close();
                log.info("success close dbr");
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
