import { NavLink } from 'react-router-dom';
import React from 'react';
import HistorySharpIcon from '@mui/icons-material/HistorySharp';
import CoPresentSharpIcon from '@mui/icons-material/CoPresentSharp';
import InfoSharpIcon from '@mui/icons-material/InfoSharp';

const NavBar = () => {
  const menuItems = [
    {
      itemPath: "/home",
      itemName: "Current",
      itemIcon: <CoPresentSharpIcon sx={{pl: .5, mt: 2.5, fontSize: '50px'}} />
    },
    {
      itemPath: "/home/history",
      itemName: "History",
      itemIcon: <HistorySharpIcon sx={{pl: .5, mt: 2.5, fontSize: '50px'}} />
    }, 
    {
      itemPath: "/home/about",
      itemName: "About",
      itemIcon: <InfoSharpIcon sx={{pl: .5, mt: 2.5, fontSize: '50px'}} />
    }
  ];

  return (
    <div>
      <div>
        {menuItems.map((item, index) => {
          return (
          <NavLink to={item.itemPath} key={index} className='link' activeclassName='active'>
            <div>{item.itemIcon}</div>
            <div>{item.itemName}</div>
          </NavLink>
        )
        })}
      </div>
    </div>
  )
};

export default NavBar;
