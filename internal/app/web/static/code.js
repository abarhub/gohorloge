

function envoyer() {
    console.log('test3');

    var checkedValue = document.querySelector('input[name="action"]:checked').value;
    console.log('valeur', checkedValue);
    if (checkedValue) {
        let param = '';
        let paramName = '';
        if (checkedValue === 'minuteur') {
            const timeControl = document.querySelector('input[type="time"]');
            if(timeControl) {
                param = timeControl.value;
                paramName = 'time';
            }
        } else if(checkedValue === 'text'){
            const textControl = document.querySelector('input[name="text-affiche"]');
            if(textControl) {
                param = textControl.value;
                paramName = 'text';
            }
        }
        fetch("/api/action/" + checkedValue + ((param !== '') ? '?'+paramName+'=' + param : ''))
            .then(response => {
                console.log('response:', response)
            })
            .catch(error => console.error(error))
    }
}

function activeChamps(action) {
    let inputMinuteur = document.getElementById('minuteur-heure');
    let inputTexteAfficher = document.getElementById('text-affiche');
    if(action==='minuteur'){
        inputMinuteur.disabled=false;
    } else {
        inputMinuteur.disabled=true;
    }
    if(action==='text'){
        inputTexteAfficher.disabled=false;
    } else {
        inputTexteAfficher.disabled=true;
    }
}

let rad = document.myForm.action;
let prev = null;
for (var i = 0; i < rad.length; i++) {
    rad[i].addEventListener('change', function() {
        (prev) ? console.log(prev.value): null;
        if (this !== prev) {
            prev = this;
        }
        console.log(this.value);
        let action=this.value;
        activeChamps(action);
    });
}

activeChamps('horloge');
