import './style.css';
import './app.css';
import { Init } from "../wailsjs/go/main/App"
import { AddElement } from "../wailsjs/go/main/App"
import { SaveTable } from "../wailsjs/go/main/App"

let table = document.getElementById("result_table");
let incr = 0;
function InitTable(){
    Init().then((result) => {
        console.log(result)
        for(const lesson of result){
            addLine(lesson);
        }
    })      
};

function addLine(lesson){
    let newRow = table.insertRow(-1);
    newRow.id = "row_" + incr;
    incr = incr + 1;
    let date = newRow.insertCell(0);
    date.textContent = lesson.date;
    let time = newRow.insertCell(1);
    time.textContent = lesson.time;
    let name = newRow.insertCell(2);
    name.textContent = lesson.name;
    let del = newRow.insertCell(3);
    del.innerHTML = '<button onclick="deleteData(this)">Delete</button>';
};

window.deleteData = function(button){
    let row = button.parentNode.parentNode;
    row.parentNode.removeChild(row);
};

window.addElement = function(){
    let data = document.getElementById("input_field");
    console.log(data.value)
    AddElement(data.value).then((result) => {
        addLine(result);
        data.value = "";
    })
};

window.saveTable = function(){
    let lessons = []
    for (var i = 0, row; row = table.rows[i]; i++) {
        let lesson = [row.cells[0].textContent,row.cells[1].textContent,row.cells[2].textContent];
        lessons[i] = lesson
        console.log(lesson)
        console.log(lessons)
    };
    SaveTable(lessons);
};


InitTable();