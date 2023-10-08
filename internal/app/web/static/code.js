console.log('test1');
console.log('test2');


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
            //.then(data => {
                //console.log(data.count)
                //console.log(data.products)
            //})
            .catch(error => console.error(error))
    }
}
