import React, { useState, useEffect, MouseEvent } from 'react'
import Terminal, { ColorMode, TerminalInput, TerminalOutput } from 'react-terminal-ui'


export const App = () => {
  const [colorMode, setColorMode] = useState(ColorMode.Dark);
  const [lineData, setLineData] = useState([
    <TerminalOutput>Read Bitch!</TerminalOutput>
  ])

  const changeColorMode = (e: MouseEvent) => {
    e.preventDefault();
    setColorMode(colorMode === ColorMode.Light ? ColorMode.Dark : ColorMode.Light);
  }

  const onInput = (input: string) => {
    let ld = [...lineData];
    ld.push(<TerminalInput>{input}</TerminalInput>);
    if(input.toLocaleLowerCase().trim() === 'clear') {
      ld = [];
    } else if(input) {
      ld.push(<TerminalInput>shit up and read!</TerminalInput>)
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
        <button className={ btnClasses.join('') } onClick={ changeColorMode }>Enable { colorMode === ColorMode.Light ? 'Dark' : 'Light'} Mode</button>
      </div>
      <Terminal
        name="Read Bitch!"
        colorMode={ colorMode }
        onInput={ onInput }
        redBtnCallback={ redBtnClick }
        yellowBtnCallback={ yellowBtnClick }
        greenBtnCallback={ greenBtnClick }
      >
        {lineData}
      </Terminal>
    </div>
  )
}


export default App
