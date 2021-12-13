require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");
import confirmDelete from "./_confirm_delete";

$(() => {
    confirmDelete.setup();
});