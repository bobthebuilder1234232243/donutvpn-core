const {  donutvpnClient } = require('./client.js');
const donutvpn = require("./donutvpn_grpc_web_pb.js");

function openConnectionPage() {
    
        $("#extension-list-container").show();
        $("#extension-page-container").hide();
        $("#connection-page").show();
        connect();
        $("#connect-button").click(async () => {
            const hsetting_request = new donutvpn.ChangeDonutSettingsRequest();
            hsetting_request.setDonutSettingsJson($("#donutvpn-settings").val());
            try{
                const hres=await donutvpnClient.changeDonutSettings(hsetting_request, {});
            }catch(err){
                $("#donutvpn-settings").val("")
                console.log(err)
            }
            
            const parse_request = new donutvpn.ParseRequest();
            parse_request.setContent($("#config-content").val());
            try{
                const pres=await donutvpnClient.parse(parse_request, {});
                if (pres.getResponseCode() !== donutvpn.ResponseCode.OK){
                    alert(pres.getMessage());
                    return
                }
                $("#config-content").val(pres.getContent());
            }catch(err){
                console.log(err)
                alert(JSON.stringify(err))
                                return
            }

            const request = new donutvpn.StartRequest();
    
            request.setConfigContent($("#config-content").val());
            request.setEnableRawConfig(false);
            try{
                const res=await donutvpnClient.start(request, {});
                console.log(res.getCoreState(),res.getMessage())
                    handleCoreStatus(res.getCoreState());
            }catch(err){
                console.log(err)
                alert(JSON.stringify(err))
                return
            }

            
        })

        $("#disconnect-button").click(async () => {
            const request = new donutvpn.Empty();
            try{
                const res=await donutvpnClient.stop(request, {});
                console.log(res.getCoreState(),res.getMessage())
                handleCoreStatus(res.getCoreState());
            }catch(err){
                console.log(err)
                alert(JSON.stringify(err))
                return
            }
        })
}


function connect(){
    const request = new donutvpn.Empty();
    const stream = donutvpnClient.coreInfoListener(request, {});
    stream.on('data', (response) => {
        console.log('Receving ',response);
        handleCoreStatus(response);
    });
    
    stream.on('error', (err) => {
        console.error('Error opening extension page:', err);
        // openExtensionPage(extensionId);
    });
    
    stream.on('end', () => {
        console.log('Stream ended');
        setTimeout(connect, 1000);
        
    });
}


function handleCoreStatus(status){
    if (status == donutvpn.CoreState.STOPPED){
        $("#connection-before-connect").show();
        $("#connection-connecting").hide();
    }else{
        $("#connection-before-connect").hide();
        $("#connection-connecting").show();
        if (status == donutvpn.CoreState.STARTING){
            $("#connection-status").text("Starting");
            $("#connection-status").css("color", "yellow");
        }else if (status == donutvpn.CoreState.STOPPING){
            $("#connection-status").text("Stopping");
            $("#connection-status").css("color", "red");
        }else if (status == donutvpn.CoreState.STARTED){
            $("#connection-status").text("Connected");
            $("#connection-status").css("color", "green");
        }
    }
}


module.exports = { openConnectionPage };