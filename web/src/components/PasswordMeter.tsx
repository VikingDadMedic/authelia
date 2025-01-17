import React, { useEffect, useState } from "react";

import { makeStyles } from "@material-ui/core";
import classnames from "classnames";
import { useTranslation } from "react-i18next";
import zxcvbn from "zxcvbn";

import { PasswordPolicyConfiguration, PasswordPolicyMode } from "@models/PasswordPolicy";

export interface Props {
    value: string;
    policy: PasswordPolicyConfiguration;
}

const PasswordMeter = function (props: Props) {
    const [progressColor] = useState(["#D32F2F", "#FF5722", "#FFEB3B", "#AFB42B", "#62D32F"]);
    const [passwordScore, setPasswordScore] = useState(0);
    const [maxScores, setMaxScores] = useState(0);
    const [feedback, setFeedback] = useState("");
    const { t: translate } = useTranslation("Portal");
    const style = makeStyles((theme) => ({
        progressBar: {
            height: "5px",
            marginTop: "2px",
            backgroundColor: "red",
            width: "50%",
            transition: "width .5s linear",
        },
    }))();

    useEffect(() => {
        const password = props.value;
        if (props.policy.mode === PasswordPolicyMode.Standard) {
            //use mode mode
            setMaxScores(4);
            if (password.length < props.policy.min_length) {
                setPasswordScore(0);
                setFeedback(
                    translate("Must be at least {{len}} characters in length", { len: props.policy.min_length }),
                );
                return;
            }
            if (props.policy.max_length !== 0 && password.length > props.policy.max_length) {
                setPasswordScore(0);
                setFeedback(
                    translate("Must not be more than {{len}} characters in length", { len: props.policy.max_length }),
                );
                return;
            }
            setFeedback("");
            let score = 1;
            let required = 0;
            let hits = 0;
            let warning = "";
            if (props.policy.require_lowercase) {
                required++;
                const hasLowercase = /[a-z]/.test(password);
                if (hasLowercase) {
                    hits++;
                } else {
                    warning += "* " + translate("Must have at least one lowercase letter") + "\n";
                }
            }

            if (props.policy.require_uppercase) {
                required++;
                const hasUppercase = /[A-Z]/.test(password);
                if (hasUppercase) {
                    hits++;
                } else {
                    warning += "* " + translate("Must have at least one UPPERCASE letter") + "\n";
                }
            }

            if (props.policy.require_number) {
                required++;
                const hasNumber = /[0-9]/.test(password);
                if (hasNumber) {
                    hits++;
                } else {
                    warning += "* " + translate("Must have at least one number") + "\n";
                }
            }

            if (props.policy.require_special) {
                required++;
                const hasSpecial = /[^0-9\w]/i.test(password);
                if (hasSpecial) {
                    hits++;
                } else {
                    warning += "* " + translate("Must have at least one special character") + "\n";
                }
            }
            score += hits > 0 ? 1 : 0;
            score += required === hits ? 1 : 0;
            if (warning !== "") {
                setFeedback(translate("The password does not meet the password policy") + ":\n" + warning);
            }
            setPasswordScore(score);
        } else if (props.policy.mode === PasswordPolicyMode.ZXCVBN) {
            //use zxcvbn mode
            setMaxScores(5);
            const { score, feedback } = zxcvbn(password);
            setFeedback(feedback.warning);
            setPasswordScore(score);
        }
    }, [props, translate]);

    return (
        <div style={{ width: "100%" }}>
            <div
                title={feedback}
                className={classnames(style.progressBar)}
                style={{
                    width: `${(passwordScore + 1) * (100 / maxScores)}%`,
                    backgroundColor: progressColor[passwordScore],
                }}
            />
        </div>
    );
};

PasswordMeter.defaultProps = {
    minLength: 0,
};

export default PasswordMeter;
