package id.my.idtecsi.inventory.api.inventory.bootstrap;

import java.sql.Connection;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class Container {
    private final DatabaseService database;

    public Container(DatabaseService database) {
        this.database = database;
    }

    @Bean
    public Connection dbw() {
        return database.getDbw();
    }

    @Bean
    public Connection dbr() {
        return database.getDbr();
    }
}
