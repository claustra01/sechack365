import styles from './devices.module.css'

export const MobileOnly = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className={styles.mobileOnly}>
      {children}
    </div>
  )
}

export const DesktopOnly = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className={styles.desktopOnly}>
      {children}
    </div>
  )
}
