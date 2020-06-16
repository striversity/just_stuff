function onLoaded() {
    console.log("onLoaded() called")
    var source = new EventSource("/sse/dashboard");

    source.onmessage = function (event) {
        console.log("OnMessage called:")
        console.dir(event);

        document.getElementById("counter").innerHTML = event.data;
    }

    source.addEventListener("tmpl", function (event) {
        console.log("got event for 'tmpl':")
        console.dir(event);

        document.getElementById("output").innerHTML = event.data;
    });
}