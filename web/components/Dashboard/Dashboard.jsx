import React from "react";

import styles from "./Dashboard.module.css";
import { Cloud } from './Cloud'


export const Dashboard = ({
  numberOfDailyMessage = 0,
  topAccountsByMessages = [],
  topMessagesByEngagements = [],
  words,
  hashtags
}) => {
  const ref = React.useRef(null)
  const [dimension, setDimension] = React.useState({ width: 0, height: 0 })
  React.useEffect(() => {
    const { offsetWidth: width } = ref.current
    setDimension({ width: width - 32, height: 320 - 32 - 28 })
  }, [])
  return (
    <main className={styles.container}>
      <h1 className={styles.title}>Dashboard</h1>
      <section ref={ref} className={styles["flex-container"]}>
        <div className={styles["card"]}>
          <h2 className={styles["card__title"]}>Number of daily messages</h2>
          <p className={styles["card__amount"]}>{numberOfDailyMessage}</p>
        </div>

        <div className={styles["card"]}>
          <h2 className={styles["card__title"]}>Top 10 accounts by messages</h2>
          <ol className={styles["list"]}>
            {topAccountsByMessages.map(account => (
              <li key={account._id} className={styles["list__item"]}>
                {account._id}
              </li>
            ))}
          </ol>
        </div>

        <div className={styles["card"]}>
          <h2 className={styles["card__title"]}>
            Top 10 messages by engagements
          </h2>
          <ol className={styles["list"]}>
            {topMessagesByEngagements.map(message => (
              <li key={message._id} className={styles["list__item"]}>
                {message.message}
              </li>
            ))}
          </ol>
        </div>

        <div className={styles["card"]}>
          <h2 className={styles["card__title"]}>Hashtag Clouds</h2>
          <Cloud {...dimension} data={hashtags} />
        </div>

        <div className={styles["card"]}>
          <h2 className={styles["card__title"]}>Word Clouds</h2>
          <Cloud {...dimension} data={words} />
        </div>

      </section>
    </main>
  );
};

export default Dashboard;
