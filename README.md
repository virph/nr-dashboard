# NR Dashboard

Program to enable fast and structurized dashboard creation.

Step:
1. Create a data csv as in [this example](https://docs.google.com/spreadsheets/d/1a0qdSLBxleFSTXAilRldUTkqr-JT5PnuVT2TpMtVJUU/edit).
2. Export the metrics data and put on the same directory of the binary as `data.csv`.
3. Execute the binary with `-title` flag as the dashboard title. Default dashboard title will be applied if no param passed.
4. Copy the JSON data written output. Can also use `./[binary file] > result.json` to keep into a file.
5. On the New Relic's Dashboards page, click the `Import Dashboard` button on the top right screen.
6. Paste the JSON data into the `Paste your JSON code` textbox.
7. Submit the dashboard by click the `Import Dashboard` button.
8. If success, the dashboard will be created with the generated title and data.