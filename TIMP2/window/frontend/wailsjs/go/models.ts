export namespace main {
	
	export class Lesson {
	    date: string;
	    time: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Lesson(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date = source["date"];
	        this.time = source["time"];
	        this.name = source["name"];
	    }
	}

}

