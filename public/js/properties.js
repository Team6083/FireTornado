function listenType(type, callbackNew, callbackModify, callbackRemove){
    
db.collection("properties").where("type", "==", 1).onSnapshot(function (snapshot) {
    snapshot.docChanges().forEach(function (change) {
        if (change.type === "added") {
            callbackNew(change.doc.data());
            console.log("New city: ", change.doc.data());
        }
        if (change.type === "modified") {
            callbackModify(change.doc.data());
            console.log("Modified city: ", change.doc.data());
        }
        if (change.type === "removed") {
            callbackRemove(change.doc.data());
            console.log("Removed city: ", change.doc.data());
        }
    });
});
}