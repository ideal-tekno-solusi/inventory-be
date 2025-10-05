package id.my.idtecsi.inventory.api.inventory.service;

public class MidtransPaymentService implements PaymentService {

    @Override
    public void ProcessPayment(double amount) {
        System.out.println("Midtrans");
        System.out.println("Amount: " + amount);
    }

}
