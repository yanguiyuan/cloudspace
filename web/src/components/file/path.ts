export class Path {
    data:string[]
    constructor() {
        this.data=[]
    }
    join(path:string):void{
        let splits=path.split("/");
        for (const s of splits) {
          if(s!=""){
              this.data=this.data.concat(s);
          }
        }
    }
    toString():string{
        let str="";
        for (const p of this.data) {
            str+=p;
            if(p!=this.data.at(this.data.length-1)){
                str+="/";
            }
        }
        return str;
    }
}
