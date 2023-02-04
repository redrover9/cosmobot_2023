//needed packages - already accomplished with package.json (?)
import { createClient } from 'pexels';
//declare constants 
const API_KEY = process.env.API_KEY;
const CLIENT = createClient(API_KEY);
//function to pull contents of text file into memory as array of words
//function to randomly select items from these arrays
//function to randomly select font
//function to create a caption
//function to get photo
//function to place caption on photo and present to user