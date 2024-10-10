<div id="createServerModal" class="modal"><div class="modal-dialog">
<div class="modal-content">
<div class="modal-header">
<h3 class="modal-title">{{.LANG.Modal.addServer.title}}</h3> <a href="#close" title="{{.LANG.Modal.control.close}}" class="close" onclick="closeModal('createServer')">X</a></div><div class="modal-body">
<div id="createServerErr"></div>
<div id="createServerSingle">
<div style="margin: 20px 0 0 0;"><p>{{.LANG.Modal.addServer.label_ip}}</p><input id="createServerIP" type="text" class="inputtext" value="" placeholder="{{.LANG.Modal.addServer.placeholder_ip}}"></div>
<div style="margin: 20px 0 0 0;"><p>{{.LANG.Modal.addServer.label_wait_time}}</p><input id="createServerWaitTime" type="text" class="inputtext" value="1800" placeholder="{{.LANG.Modal.addServer.placeholder_wait_time}}"></div>
<div style="margin: 20px 0 0 0;"><p>{{.LANG.Modal.addServer.label_usename}}</p><input id="createServerUsername" type="text" class="inputtext" value="root"></div>
<div style="margin: 20px 0 0 0;"><p>{{.LANG.Modal.addServer.label_password}}</p><input id="createServerPassword" type="text" class="inputtext" value="1Htaht;bhfnjh"></div>
<div style="margin: 20px 0 0 0;"><p>{{.LANG.Modal.addServer.label_ssh_key}}</p><textarea id="createServerKeySsh" class="inputtext" rows="10" cols="50"></textarea></div>
<div style="margin: 20px 0 0 0;"><input type="checkbox" id="useProxy"> <label for="useProxy">_{{.LANG.Modal.addServer.label_use_proxy}}</label></div>

</div>

<div style="text-align: center">
<input type="button" value="{{.LANG.Modal.addServer.label_create}}" onclick="createServer()" id="createServerButton" class="sendbutton">
</div>
</div>
</div>
</div>

</div>

