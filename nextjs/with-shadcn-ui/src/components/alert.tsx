'use client';

import { Button } from "./ui/button";

const AlertButton = () => {
  const handleClick = () => alert("test");
  return (
    <>
      <Button onClick={handleClick}>test</Button>
    </>
  )
}

export default AlertButton;
