import React, { useState, useEffect, MouseEvent } from 'react'
import Terminal, { ColorMode, TerminalInput, TerminalOutput } from 'react-terminal-ui'
import axios, { AxiosProgressEvent, AxiosResponse } from 'axios';

type Book = {
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date | null;
  title: string;
  description: string;
  rate: number;
  categoryID: number;
  userID: number;
}

type Category = {
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date | null;
  name: string;
  userID: number;
  books: Book[] | null;
  parentID: number | null;
  subcategories: Category[] | null;
}

type User = {
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date | null;
  ID: number;
  username: string;
  email: string;
  password: string;
  books: Book[] | null;
  categories: Category[] | null;
}

export const App = () => {
  const [colorMode, setColorMode] = useState(ColorMode.Dark);
  const [lineData, setLineData] = useState([
    <TerminalOutput>Read Bitch!</TerminalOutput>
  ])
  const [awaitingUsername, setAwaitingUsername] = useState<boolean>(false)
  const [awaitingEmail, setAwaitingEmail] = useState<boolean>(false)
  const [awaitingPassword, setAwaitingPassword] = useState<boolean>(false)
  const [loginFlag, setLoginFlag] = useState<boolean>(false)

  const [email, setEmail] = useState<string>('')
  const [username, setUsername] = useState<string>('')
  const [categoryID, setCategoryID] = useState<number>(0); // use to locate the current directy. when is 0, it means ur in root directy

  const [token, setToken] = useState<string>(localStorage.getItem('token') ?? '')
  const [user, setUser] = useState<User | undefined>(JSON.parse(localStorage.getItem('user') ?? "{}"))


  const changeColorMode = (e: MouseEvent) => {
    e.preventDefault();
    setColorMode(colorMode === ColorMode.Light ? ColorMode.Dark : ColorMode.Light);
  }

  const addBook = async (title: string, description: string, rate: number) => {
    const category_id = categoryID == 0 ? null : categoryID;
    if (user == undefined) {
      setLineData(prev => [...prev, <TerminalOutput>You are not login. use login command to login</TerminalOutput>])
      return;
    }
    try {
      const res: AxiosResponse = await axios.post(`http://localhostL8080/add-book?category_id=${categoryID}`, {
        title,
        description,
        rate,
        category_id,
        userID: user?.ID
      })
      setLineData(prev => [...prev, <TerminalOutput>add book ok</TerminalOutput>])
    } catch (err) {
      setLineData(prev => [...prev, <TerminalOutput>Something went Wrong!</TerminalOutput>])
    }
  }

  const addCategory = async (categoryName: string) => {

  }

  const loginUser = async (email: string, password: string) => {
    setLineData(prev => [...prev, <TerminalOutput>waiting for server...</TerminalOutput>])
    try {
      const res: AxiosResponse = await axios.post("http://localhost:8080/login", {
        email,
        password
      })
      if (res.data.token) {
        localStorage.setItem("token", res.data.token);
        localStorage.setItem("user", JSON.stringify(res.data.user));
        setLineData(prev => [...prev, <TerminalOutput>login ok</TerminalOutput>])
      }
    } catch (err) {
      setLineData(prev => [...prev, <TerminalOutput>Something went wrong!</TerminalOutput>])
    }
  }

  const registerUser = async (email: string, password: string, username: string) => {
    setLineData(prev => [...prev, <TerminalOutput>waiting for server...</TerminalOutput>])
    try {
      await axios.post("http://localhost:8080/register", {
        email,
        password,
        username
      })
      setLineData(prev => [...prev, <TerminalOutput>register ok</TerminalOutput>])
    } catch (err) {
      setLineData(prev => [...prev, <TerminalOutput>Something went wrong!</TerminalOutput>])
    }
  }

  const onInput = (input: string) => {
    let ld = [...lineData];
    ld.push(<TerminalInput>{input}</TerminalInput>);

    if (input.toLocaleLowerCase().trim() === 'clear') {
      ld = [];
    } else if (input.toLocaleLowerCase().trim() === 'login') { // login
      if (localStorage.getItem("token") != undefined) { // if you have token in localstorage (you are already login)
        setLineData(prev => [...prev, <TerminalInput>You have already login</TerminalInput>])
        return;
      } else {
        ld.push(<TerminalInput>Enter email</TerminalInput>)
        setLoginFlag(true);
        setAwaitingEmail(true); // go to #1
      }
    } else if (input.toLocaleLowerCase().trim() === 'logout') { // logout
      if (localStorage.getItem("token") != undefined) { // if you have token in localstorage (you are already login)
        localStorage.removeItem("token");
        localStorage.removeItem("user");
        setLineData(prev => [...prev, <TerminalInput>logout ok</TerminalInput>])
        return;
      } else {
        setLineData(prev => [...prev, <TerminalInput>You are not login. use login command to login</TerminalInput>])
        return;
      }
    } else if (input.toLocaleLowerCase().trim() === 'register') { // register
      if (localStorage.getItem("token") != undefined) { // if you have token in localstorage (you are already login)
        setLineData(prev => [...prev, <TerminalInput>You have already login</TerminalInput>])
      } else {
        ld.push(<TerminalInput>Enter username</TerminalInput>)
        setAwaitingUsername(true); // go to #1
      }
    } else if (input.toLocaleLowerCase().trim() === 'addbook') {

    }

    // #1
    if (awaitingUsername) {
      setUsername(input);
      setAwaitingUsername(false);
      setAwaitingEmail(true);
      ld.push(<TerminalInput>Enter email</TerminalInput>)
    } else if (awaitingEmail) {
      setEmail(input);
      setAwaitingEmail(false);
      setAwaitingPassword(true);
      ld.push(<TerminalInput>Enter password</TerminalInput>)
    } else if (awaitingPassword) {
      if (!loginFlag) {
        registerUser(email, input, username);
        setAwaitingPassword(false);
      } else {
        loginUser(email, input);
        setAwaitingPassword(false);
      }
    }

    setLineData(ld);

  }

  const redBtnClick = () => {
    console.log("Clicked the red button.");
  }

  const yellowBtnClick = () => {
    console.log("Clicked the yellow button.");
  }

  const greenBtnClick = () => {
    console.log("Clicked the green button.");
  }

  const btnClasses = ['btn'];
  if (colorMode === ColorMode.Light) {
    btnClasses.push('btn-dark');
  } else {
    btnClasses.push('btn-light');
  }

  return (
    <div className="container">
      <div className="d-flex flex-row-reverse p-2">
        <button className={btnClasses.join('')} onClick={changeColorMode}>Enable {colorMode === ColorMode.Light ? 'Dark' : 'Light'} Mode</button>
      </div>
      <Terminal
        name="Read Bitch!"
        colorMode={colorMode}
        onInput={onInput}
        redBtnCallback={redBtnClick}
        yellowBtnCallback={yellowBtnClick}
        greenBtnCallback={greenBtnClick}
      >
        {lineData}
      </Terminal>
    </div>
  )
}


export default App
