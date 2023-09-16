import axios from 'axios'
const cf=axios.create({
    baseURL:'http://127.0.0.1:8081/api/cf'
})
const js=axios.create({
    baseURL:"http://127.0.0.1:8080/nt"
})

// cf.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
// cf.defaults.headers.put['Content-Type'] = 'application/x-www-form-urlencoded';
let token=localStorage.getItem("token")
if(token==null){
    token=""
}else{
    token=token.substring(1,token.length-1)
}
let tokenHeader="Bearer ".concat(token)
cf.defaults.headers.common['Authorization'] = tokenHeader
console.log(tokenHeader)
export {js,cf}