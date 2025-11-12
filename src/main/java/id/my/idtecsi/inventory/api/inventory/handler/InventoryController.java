package id.my.idtecsi.inventory.api.inventory.handler;

import java.sql.Connection;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import id.my.idtecsi.inventory.api.inventory.bootstrap.DatabaseService;
import id.my.idtecsi.inventory.api.inventory.entity.*;
import jakarta.validation.Valid;

@RestController
@RequestMapping("/api/v1")
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
    public ResponseEntity<DomainInventoryResponse> Inventory(@Valid @ModelAttribute DomainInventoryRequest req) {
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

        return new ResponseEntity<>(null, HttpStatus.OK);
    }
}
