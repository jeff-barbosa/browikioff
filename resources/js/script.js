function search(str) {
  if (str.length == 0) {
    document.getElementById("bwo-searchbox").innerHTML = "";
    document.getElementById("bwo-searchbox").style.display = "none"
    return;
  }

  xmlhttp = new XMLHttpRequest();

  xmlhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
      document.getElementById("bwo-searchbox").innerHTML = this.responseText;
      document.getElementById("bwo-searchbox").style.display = "block"
    }
  }

  xmlhttp.open("GET", "http://localhost:8080/ajax?q="+str, true);
  xmlhttp.send();
}