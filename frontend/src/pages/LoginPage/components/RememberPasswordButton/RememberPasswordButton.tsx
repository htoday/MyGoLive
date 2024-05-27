import styles from "./RememberPasswordButton.module.less"

export function RememberPasswordButton(props: {
    onClick: () => void,
    state: boolean;
}) {

    return (
        <div className={styles.row}>
            <div
                className={props.state ? styles.checkbox_checked : styles.checkbox_unchecked}
                onClick={props.onClick}
            />
            <h1>记住密码</h1>
        </div>
    )
}