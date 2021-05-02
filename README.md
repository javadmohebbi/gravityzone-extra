# Bitdefender GravityZone Extra
This repo intent to add extra reports to Bitdefender GravityZone

List of available reports:
- [x] [Application Inventory Report](#application-inventory-report): First, you must run **inventory task** on your endpoits. Then, this app will help you get a **CSV** comma delimited report on the app. inventory.


## Application Inventory Report
To get application inventory report, do the following steps on your **GravityZone Appliance**:
```
wget -O /home/bdadmin/get-inventory-report  https://github.com/javadmohebbi/gravityzone-extra/raw/main/build/get-inventory-report

chmod +x -v /home/bdadmin/get-inventory-report

/home/bdadmin/get-inventory-report -mongo-pass {MONGO_DB_PASS}

#*** replace your mongo database password with {MONGO_DB_PASS}

```
- *** **Mongo DB password** must be encoded using **percent encoding**. Use [this website](https://www.url-encode-decode.com/) to encode your password.v
- This tiny app will extract needed information and will create two **csv** file:
  - **apps.csv**: A comma delimited CSV file with the further template:
    - Application Name, Application Group, Hash, Version, Discover Date, NumberOfEndpoints
  - **app_details.csv**: A comma delimited CSV file with the further template:
    - Endpoint Name, Endpoint OS, Application Name, Application Group, Hash, Version, Discover Date
- By default mentioned csv files will placed in **/home/bdadmin**




### Need more custom report?
Just send a request to [javad [at] openintelligence24 [dot] com](mailto:javad@openintelligence24.com)