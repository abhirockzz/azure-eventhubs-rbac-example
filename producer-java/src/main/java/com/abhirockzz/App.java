package com.abhirockzz;

import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.messaging.eventhubs.*;

public class App {
    static String ehConnectionStringEnvVarName = "EVENTHUB_CONNECTION_STRING";

    public static void main(final String[] args) throws Exception {

        String eventhubsNamespace = System.getenv("EVENTHUBS_NAMESPACE");
        String eventhubName = System.getenv("EVENTHUB_NAME");

        // create a producer using the namespace connection string and event hub name
        EventHubProducerClient producer = new EventHubClientBuilder()
                .credential(eventhubsNamespace, eventhubName, new DefaultAzureCredentialBuilder().build())
                .buildProducerClient();

        EventDataBatch batch = producer.createBatch();
        for (int i = 1; i <= 10; i++) {
            System.out.println("adding event " + i);
            batch.tryAdd(new EventData("event-" + i));
        }
        producer.send(batch);
        producer.close();
    }
}