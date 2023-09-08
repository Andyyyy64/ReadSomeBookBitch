import React, { useState, useEffect, MouseEvent } from 'react'
import Terminal, { ColorMode, TerminalInput, TerminalOutput } from 'react-terminal-ui'
import axios, { AxiosResponse } from 'axios';

export const App = () => {
  const [colorMode, setColorMode] = useState(ColorMode.Dark);
  const [lineData, setLineData] = useState([
    <TerminalOutput>Read Bitch!</TerminalOutput>
  ])
  const [awaitingEmail, setAwaitingEmail] = useState<Boolean>(false)
  const [awaitingPassword, setAwaitingPassword] = useState<Boolean>(false)
  const [email, setEmail] = useState<string>('')

  const changeColorMode = (e: MouseEvent) => {
    e.preventDefault();
    setColorMode(colorMode === ColorMode.Light ? ColorMode.Dark : ColorMode.Light);
  }

  const loginUser = async (email: string, password: string) => {
    try {
      const res: AxiosResponse = await axios.post("http://localhost:8080/login", {
        email,
        password
      });
      if (res.data.token) {
        localStorage.setItem("token", res.data.token);
        setLineData(prev => [...prev, <TerminalOutput>login ok</TerminalOutput>])
      }
    } catch (err) {
      setLineData(prev => [...prev, <TerminalOutput>Something went wrong!</TerminalOutput>])
    }
  }

  const onInput = (input: string) => {
    let ld = [...lineData];
    ld.push(<TerminalInput>{input}</TerminalInput>);

    if (input.toLocaleLowerCase().trim() === 'clear') {
      ld = [];
    } else if (input.toLocaleLowerCase().trim() === 'login') {
      ld.push(<TerminalInput>Enter email</TerminalInput>)
      setAwaitingEmail(true); // go to login
    }

    // login
    if (awaitingEmail) {
      setEmail(input);
      setAwaitingEmail(false);
      setAwaitingPassword(true);
      ld.push(<TerminalInput>Enter password</TerminalInput>)
    } else if (awaitingPassword) {
      loginUser(email, input);
      setAwaitingPassword(false);
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
