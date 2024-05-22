# Postgres testapp

En veldig enkel app som ble brukt for å teste responstid på Postgres-databaser når Knada kjørte i Belgia, mens resten av NAV kjørte i Finland.
Viste seg at alle kall mot våre belgiske databaser fra et belgisk Kubernetes cluster gikk via Finland på grunn av peering.
Denne appen var veldig nyttig for å enkelt kunne kjøre tester for å få god data.

## Utvikling

Bygges manuelt, og rulles manuelt ut i NAIS.
