import React, { useState, useEffect } from 'react';
import { Container } from 'react-bootstrap';

function Countdown({ targetDate }) {
  const calculateTimeLeft = () => {
    const difference = +new Date(targetDate) - +new Date();
    let timeLeft = {};

    if (difference > 0) {
      timeLeft = {
        days: Math.floor(difference / (1000 * 60 * 60 * 24)),
        hours: Math.floor((difference / (1000 * 60 * 60)) % 24),
        minutes: Math.floor((difference / 1000 / 60) % 60),
        seconds: Math.floor((difference / 1000) % 60),
      };
    }

    return timeLeft;
  };

  const [timeLeft, setTimeLeft] = useState(calculateTimeLeft());

  useEffect(() => {
    const timer = setTimeout(() => {
      setTimeLeft(calculateTimeLeft());
    }, 1000);

    return () => clearTimeout(timer);
  }, [timeLeft, targetDate]);

  return (
    <Container>
    <div className="countdown d-flex flex-wrap justify-content-between text-center">
      <div className="countdown-holder d-flex flex-column justify-content-center align-items-center mx-2">
        <div className='dday'>{timeLeft.days}</div>
        <label>Days</label>
      </div>
      <div className="countdown-holder d-flex flex-column justify-content-center align-items-center mx-2">
        <div className='dhour'>{timeLeft.hours}</div>
        <label>Hours</label>
      </div>
      <div className="dmin countdown-holder d-flex flex-column justify-content-center align-items-center mx-2">
        <div className='dmin'>{timeLeft.minutes}</div>
        <label>Minutes</label>
      </div>
      <div className="countdown-holder d-flex flex-column justify-content-center align-items-center mx-2">
        <div className='dsec' >{timeLeft.seconds}</div>
        <label>Seconds</label>
      </div>
    </div>
    </Container>
  );
}

export default Countdown;
