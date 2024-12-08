import styles from './admin-home.module.css';


const AdminHome = () => {
  	return (
    		<div className={styles.adminHome}>
      			<div className={styles.rectangleParent}>
        				<div className={styles.rectangle} />
        				<div className={styles.thereCurrentlyAre}>There currently are:</div>
        				<div className={styles.heyAdmin}>Hey Admin!</div>
        				<div className={styles.groupParent}>
          					<div className={styles.groupWrapper}>
            						<div className={styles.groupWrapper}>
              							<img className={styles.rectangleIcon} alt="" src="Rectangle.svg" />
              							<div className={styles.resolvedTickets}>Resolved Tickets</div>
            						</div>
          					</div>
          					<div className={styles.div}>56</div>
        				</div>
        				<div className={styles.groupContainer}>
          					<div className={styles.groupWrapper}>
            						<img className={styles.rectangleIcon} alt="" src="Rectangle.svg" />
            						<div className={styles.resolvedTickets}>Unresolved Tickets</div>
          					</div>
          					<div className={styles.div1}>26</div>
        				</div>
        				<div className={styles.groupDiv}>
          					<div className={styles.groupWrapper}>
            						<img className={styles.rectangleIcon} alt="" src="Rectangle.svg" />
            						<div className={styles.assignedTickets}>Assigned Tickets</div>
          					</div>
          					<div className={styles.div1}>56</div>
        				</div>
        				<div className={styles.groupParent1}>
          					<div className={styles.groupWrapper}>
            						<img className={styles.rectangleIcon} alt="" src="Rectangle.svg" />
            						<div className={styles.unassignedTickets}>Unassigned Tickets</div>
          					</div>
          					<div className={styles.div3}>20</div>
        				</div>
      			</div>
    		</div>);
};

export default AdminHome;
