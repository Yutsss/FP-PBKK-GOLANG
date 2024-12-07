// import React from 'react';
// import "./checkTicket.css";

// const UserCheckTicketsLoggedIn = () => {
//   const ticketData = [
//     { status: 'Unresolved' },
//     { status: 'Unresolved' },
//     { status: 'Resolved' },
//     { status: 'Unresolved' },
//   ];

//   return (
//     <div className="userCheckTicketsLoggedIn" style={styles.container}>
//       <Header />
//       {ticketData.map((ticket, index) => (
//         <Ticket key={index} status={ticket.status} />
//       ))}
//     </div>
//   );
// };


// const Ticket = ({status}) => (
//     <div style={styles.ticketContainer}>
//         <div style={styles.ticketRectangle}></div>
//         <div style={styles.ticketText}>Ticket Subject: Category: Body: Technician:</div>
//         <div style={styles.buttonContainer}>
//              <button style={{...styles.button, ...(status === 'Resolved' ? styles.resolvedButton : styles.unresolvedButton)}}>
//                 {status}
//             </button>
//         </div>
//     </div>
// );

// export default UserCheckTicketsLoggedIn;