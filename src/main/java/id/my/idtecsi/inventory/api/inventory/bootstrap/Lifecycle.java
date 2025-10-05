package id.my.idtecsi.inventory.api.inventory.bootstrap;

import org.springframework.beans.factory.DisposableBean;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

import id.my.idtecsi.inventory.api.inventory.InventoryApplication;

public class Lifecycle implements DisposableBean {
    private ConfigurableApplicationContext applicationContext;

    @Override
    public void destroy() throws Exception {
        applicationContext = new AnnotationConfigApplicationContext(
                InventoryApplication.class);

        DatabaseService conn = applicationContext.getBean(PostgresDatabaseService.class);

        applicationContext.close();
        conn.Close();
    }

}
