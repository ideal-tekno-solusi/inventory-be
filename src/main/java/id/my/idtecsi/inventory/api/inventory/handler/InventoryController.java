package id.my.idtecsi.inventory.api.inventory.handler;

import java.sql.SQLException;

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
import id.my.idtecsi.inventory.api.inventory.entity.PageInformation;
import id.my.idtecsi.inventory.api.inventory.model.*;
import id.my.idtecsi.inventory.api.inventory.repository.InventoryRepository;
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
    public ResponseEntity<InventoryResponse> Inventory(@Valid @ModelAttribute InventoryRequest req)
            throws SQLException {
        // TODO:lanjutin buat logic dari verifikasi req sampe return disini
        InventoryRepository inventoryRepository = new InventoryRepository(dbr, dbw);

        PageInformation pageInfo = inventoryRepository.getInventoryPageInformation(req.getCategoryId(),
                req.getBranchId(), req.getLimit());

        // TODO: lanjut untuk query fetch data inventorynya

        return new ResponseEntity<>(null, HttpStatus.OK);
    }
}
