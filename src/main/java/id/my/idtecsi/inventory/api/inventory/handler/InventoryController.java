package id.my.idtecsi.inventory.api.inventory.handler;

import java.sql.Connection;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import id.my.idtecsi.inventory.api.inventory.bootstrap.PostgresDatabaseService;

@RestController
@RequestMapping("/v1")
public class InventoryController {
    private Connection dbw;
    private Connection dbr;
    private static final Logger log = LoggerFactory.getLogger(PostgresDatabaseService.class);

    public InventoryController(Connection dbw, Connection dbr) {
        this.dbw = dbw;
        this.dbr = dbr;
    }

    @GetMapping("/inventory")
    public void Inventory() {
        // TODO:lanjutin buat logic dari verifikasi req sampe return disini
    }
}
