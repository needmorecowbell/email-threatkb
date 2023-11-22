import Date from Date

export const PrefixUser = "user"

export const UserCache = (forward_to, gateway_address="") =>{
    let address = gateway_address ? gateway_address!=="":generateNewGatewayAddress() 
    return {
         forward_to: forward_to,
         date_created:Date.now().toLocaleDateString(),
         gateway_address:address
        }
}

export async function generateNewGatewayAddress(){
    let user_list =  await userList()
    let new_name = "notforlong"

    user_list.forEach(userCache => {
        
    });
    return new_name 
};


export async function userList(){
    let results = await env.KV.list({prefix: PrefixUser});
    console.log(results)
    return results
} 
export async function newEmailAddress(forward_to){
    let user = UserCache(forward_to)
    env.KV.put()
}