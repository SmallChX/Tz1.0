import React from "react";
import HeroContent from "./Hero";
import "./style.css"
import LineupArtists from "./Timeline";
import HomepageNextEvents from "./EventLibrary";
import Introduction from "./Introduction";
import Footer from "../../components/Footer";
import Countdown from "./Countdown";


function HomePage() {
    return (
      <div>
        <HeroContent />  
        <div className="content-section pt-5">
          <Countdown targetDate="03/30/2024" />
          <Introduction />
          <LineupArtists />
          <HomepageNextEvents />
        </div>
        <Footer />
      </div>
    )
}

export default HomePage;