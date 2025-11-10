package id.my.idtecsi.inventory.api.inventory.handler;

import java.sql.Connection;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import id.my.idtecsi.inventory.api.inventory.bootstrap.PostgresDatabaseReadService;
import id.my.idtecsi.inventory.api.inventory.bootstrap.PostgresDatabaseWriteService;

@RestController
@RequestMapping("/v1")
public class InventoryController {
    private PostgresDatabaseReadService postgresDatabaseReadService;
    private PostgresDatabaseWriteService PostgresDatabaseWriteService;
    private static final Logger log = LoggerFactory.getLogger(InventoryController.class);

    public InventoryController(PostgresDatabaseReadService postgresDatabaseReadService,
            PostgresDatabaseWriteService PostgresDatabaseWriteService) {
        this.postgresDatabaseReadService = postgresDatabaseReadService;
        this.PostgresDatabaseWriteService = PostgresDatabaseWriteService;
    }

    @GetMapping("/inventory")
    public void Inventory() {
        // TODO:lanjutin buat logic dari verifikasi req sampe return disini
        try {
            Connection dbr = this.postgresDatabaseReadService.getDb();
            Connection dbw = this.PostgresDatabaseWriteService.getDb();

            log.info("dbr is closed: {}", dbr.isClosed());
            log.info("dbw is closed: {}", dbw.isClosed());
        } catch (Exception ex) {
            ex.printStackTrace();
            log.error("error on inventory handler with error: ", ex);
        }
    }
}
