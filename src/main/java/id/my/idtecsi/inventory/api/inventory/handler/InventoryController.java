package id.my.idtecsi.inventory.api.inventory.handler;

import java.sql.Connection;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import id.my.idtecsi.inventory.api.inventory.bootstrap.DatabaseService;

@RestController
@RequestMapping("/v1")
public class InventoryController {
    private DatabaseService dbr;
    private DatabaseService dbw;
    private static final Logger log = LoggerFactory.getLogger(InventoryController.class);

    public InventoryController(@Qualifier("postgresDatabaseReadService") DatabaseService dbr,
            @Qualifier("postgresDatabaseWriteService") DatabaseService dbw) {
        this.dbr = dbr;
        this.dbw = dbw;
    }

    @GetMapping("/inventory")
    public void Inventory() {
        // TODO:lanjutin buat logic dari verifikasi req sampe return disini
        try {
            Connection dbr = this.dbr.getDb();
            Connection dbw = this.dbw.getDb();

            log.info("dbr is closed: {}", dbr.isClosed());
            log.info("dbw is closed: {}", dbw.isClosed());
        } catch (Exception ex) {
            ex.printStackTrace();
            log.error("error on inventory handler with error: ", ex);
        }
    }
}
