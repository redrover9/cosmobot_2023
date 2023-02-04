//needed packages - already accomplished with package.json (?)
import './node_modules/pexels/dist/main.js';
//declare constants 
const API_KEY = process.env.API_KEY;
const CLIENT = createClient(API_KEY);
//test
CLIENT.photos.show({ id: 2014422 }).then(photo => {console.log(photo)});
//function to pull contents of text file into memory as array of words
//function to randomly select items from these arrays
//function to randomly select font
//function to create a caption
//function to get photo
//function to place caption on photo and present to user