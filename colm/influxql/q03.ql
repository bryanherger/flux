# Alternate expression tree from which we extract time.

SELECT usage_user FROM cpu
	WHERE ( cpu = 'cpu0' OR cpu = 'cpu1' ) AND ( ( ( ( time > -2m ) AND ( time < -1m ) ) ) );
